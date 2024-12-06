{ pkgs, lib, config, inputs, ... }:

{
    # https://devenv.sh/basics/
    env.GREET = "time for Advent of Code 2024";

    # https://devenv.sh/packages/
    packages = [];

    # https://devenv.sh/languages/
    languages.go.enable = true;

    # https://devenv.sh/processes/
    # processes.cargo-watch.exec = "cargo-watch";


    # https://devenv.sh/scripts/
    scripts.welcome.exec = ''
    echo Hello,$GREET
    '';

    enterShell = ''
   welcome 
    '';

    # https://devenv.sh/tasks/
    # tasks = {
    #   "myproj:setup".exec = "mytool build";
    #   "devenv:enterShell".after = [ "myproj:setup" ];
    # };

    # https://devenv.sh/tests/
    enterTest = ''
    echo "Running tests"
    '';

    # https://devenv.sh/pre-commit-hooks/
    # pre-commit.hooks.shellcheck.enable = true;

    # See full reference at https://devenv.sh/reference/options/
}
