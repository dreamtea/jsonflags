jsonflags
=========

A very simple Go package to read flags from a JSON object in addition to the command line. 
Define your flags as normal using the built-in `flag` package and call `jsonflags.Parse()` 
instead of `flag.Parse()`.

The application should define a flag called `"config"`, from which `jsonflags` will take the
name of the JSON config file to read. The default value for this flag provides a path to a 
config file load if the user does not provide a `-config` argument on the command line. In 
the case that the default path is used, `jsonflags` will not return an error if the file 
does not exist.

If the application does not define a `"config"` flag, `jsonflags.Parse()` behaves 
identically to `flag.Parse()`.

Arguments supplied on the command line override those defined by the JSON config, which in
turn override the default values.
