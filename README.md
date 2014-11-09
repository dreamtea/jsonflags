jsonflags
=========

A very simple Go package to read flags from a JSON object. Define your flags as normal using 
the built-in `flag` package, then call

    jsonflags.Parse()

instead of `flag.Parse()`.

The path of the JSON file to read is taken from a string flag called `"config"` which should 
be defined by the application, like this:

    flag.String("config", "config.json", "path to JSON config file")

If there is no `"config"` flag, `jsonflags.Parse()` behaves identically to `flag.Parse()`.

If `"config"` has a non-empty default value, it specifies a file to load if the program is 
not invoked with a `-config myconfig.json` argument. When this default path is used, 
`jsonflags` does not return an error if the file is missing.

Arguments supplied on the command line override those defined by the JSON config, which in 
turn override the defaults.
