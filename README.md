# cfgmaker
    A tiny tool, which can replace macro in a file by environment variable named the same as macro.

# format macro
    The macro is included by '%% %%', eg :
        %%ENV_DB_PWD%%
	%%ENV_XXXXXX%%

    And ensure the environment variable has value, such as :
        export ENV_DB_PWD=xxxxxxx
	export ENV_XXXXXX="123456"

# usage
    ./cfgmaker -i <input-file> -o <out-file>
