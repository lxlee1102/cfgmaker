# cfgmaker
    A tiny tool, which can replace macro in a file by environment variable named the same as macro.

# format macro
    The macro is included by '%% %%' in the input file, eg:
        %%ENV_DB_GMS%%
        %%MAX_CONN%%
        %%ENV_MCSAPI%%

    And ensure the environment variable has value, such as:
        export ENV_DB_GMS="root@127.0.0.1:3306"
        export MAX_CONN=100
        export ENV_MCSAPI=http://api.mcscon:8080

# build
    git clone https://github.com/lxlee1102/cfgmaker.git
    cd cfgmaker/
    go get
    go build

# usage
    ./cfgmaker -i <input-file> -o <out-file>
