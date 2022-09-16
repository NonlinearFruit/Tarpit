package smc;

import com.cleancoder.args.Args;
import com.cleancoder.args.ArgsException;
import smc.generators.nestedSwitchCaseGenerator.NSCGenerator;
import smc.implementers.JavaNestedSwitchCaseImplementer;
import smc.lexer.Lexer;
import smc.optimizer.Optimizer;
import smc.parser.FsmSyntax;
import smc.parser.Parser;
import smc.parser.SyntaxBuilder;
import smc.semanticAnalyzer.AbstractSyntaxTree;
import smc.semanticAnalyzer.SemanticAnalyzer;

import java.io.IOException;
import java.nio.file.FileSystems;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;

import static smc.parser.ParserEvent.EOF;

public class SMC {
  public static void main(String[] args) throws IOException {
    Args argParser;
    String argSchema = "a,p*,o*";
    try {
      argParser = new Args(argSchema, args);
      new SmcCompiler(args, argParser).run();
    } catch (ArgsException e) {
      System.out.println("usage: " + argSchema + " file");
      System.out.println(e.errorMessage());
      System.exit(0);
    }
  }

  private static class SmcCompiler {
    private String[] args;
    private Args argParser;
    private String javaPackage = null;
    private String outputDirectory = null;

    public SmcCompiler(String[] args, Args argParser) {
      this.args = args;
      this.argParser = argParser;
    }

    public void run() throws IOException {
      if (argParser.has('p'))
        javaPackage = argParser.getString('p');
      if (argParser.has('o'))
        outputDirectory = argParser.getString('o');

      String fileName = args[argParser.nextArgument()];
      String smContent = new String(Files.readAllBytes(Paths.get(fileName)));

      SyntaxBuilder syntaxBuilder = new SyntaxBuilder();
      Parser parser = new Parser(syntaxBuilder);
      Lexer lexer = new Lexer(parser);
      SemanticAnalyzer analyzer = new SemanticAnalyzer();
      Optimizer optimizer = new Optimizer();
      NSCGenerator generator = new NSCGenerator();

      lexer.lex(smContent);
      parser.handleEvent(EOF, -1, -1);

      FsmSyntax fsm = syntaxBuilder.getFsm();
      int syntaxErrorCount = fsm.errors.size();

      System.out.println(String.format("Compiled with %d syntax error%s.", syntaxErrorCount, (syntaxErrorCount == 1 ? "" : "s")));

      for (FsmSyntax.SyntaxError error : fsm.errors)
        System.out.println(error.toString());

      if (syntaxErrorCount == 0) {
        AbstractSyntaxTree ast = analyzer.analyze(fsm);
        smc.StateMachine stateMachine = optimizer.optimize(ast);

        JavaNestedSwitchCaseImplementer implementer = new JavaNestedSwitchCaseImplementer(javaPackage);
        generator.generate(stateMachine).accept(implementer);

        String outputFileName = stateMachine.header.fsm + ".java";

        Path outputPath;
        if (outputDirectory == null)
          outputPath = FileSystems.getDefault().getPath(outputFileName);
        else
          outputPath = FileSystems.getDefault().getPath(outputDirectory, outputFileName);

        Files.write(outputPath, implementer.getOutput().getBytes());
      }
    }
  }
}
