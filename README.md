# Sting of the Viper (with subcommands)

This fork of [carolynvs/stingoftheviper](https://github.com/carolynvs/stingoftheviper) shows an example on how to integrate flags, environment variables and such with different subcommands.

Please have a look at her awesome blog post about this: [Sting of the Viper: Getting Cobra and Viper to work together](https://carolynvanslyck.com/blog/2020/08/sting-of-the-viper/)

# Differences from the original app

The idea was to have a simple app with also a subcommand, so there is the `stingoftheviper` command, and the `stingoftheviper sting` subcommand.

The `main.go` and `flags.go` files have basically the same original content, I've moved the root and sting commands into their own file, so that it was clearer to see the differences.

The `config.go` file contains the struct that will hold the configuration, and it also contains the default values.

```go
type Config struct {
	Name   string
	Number int

	StingConfig StingConfig
}

type StingConfig struct {
	Name string
}
```

Before initializing the rootCmd you have to initialize your empty Config, and bind it to the flags with the `bindRootFlags` func:

```go
func bindRootFlags(flags, persistentFlags *pflag.FlagSet, config *Config) {
	flags.StringVarP(&config.Name, "name", "n", config.Name, "What's your name?")

	// this flag will be persisted trough the sub-commands
	persistentFlags.IntVarP(&config.Number, "number", "c", config.Number, "Which is your favorite number?")
}
```

I've extended the original work with also the persistentFlags.

Then, when creating the stingCmd just pass to it the Config:

```go
rootCmd.AddCommand(NewStingCommand(&config))
```

so that it will be possible for it to do the same, with the `bindStingFlags` func:

```go
func bindStingFlags(flags, persistentFlags *pflag.FlagSet, config *StingConfig) {
	flags.StringVarP(&config.Name, "name", "n", config.Name, "Who do you want to sting?")
}
```

Note: this is just a way to do this. :)


# Try it out

First grab the source code with go get or by cloning it. Change into the directory of this repository.

```
go get github.com/carolynvs/stingoftheviper
# or
git clone https://github.com/carolynvs/stingoftheviper.git
cd stingoftheviper/
```

Now let's build the CLI (stingoftheviper) and make sure everything is still working:

```
go build .
go test ./...
```

We are now ready to try out a few scenarios to test out the precedence order. First let's run it with no flags
or environment variables.

```console
$ ./stingoftheviper
Your favorite color is: blue
The magic number is: 7
```

If you take a peek at the config file, you will see that only the favorite-color was set there. So favorite-color
got its value from the config file, while magic number got its value from the flag's default value set in main.go.
So the lowest precedence is the flag default, followed by the config file.

Let's try setting an environment variable.

```console
$ STING_FAVORITE_COLOR=purple ./stingoftheviper
Your favorite color is: purple
The magic number is: 7
```

There's two interesting things going on here. One is that the environment variable has higher precedence than the
config file value obviously. The other is that not only was the environment variable automatically bound to
the flag, but we handled swapping the dashes for underscores in the binding (which isn't done for us in the 
library, you have to do that yourself).

To finish things off, let's actually use a flag.

```console
$ ./stingoftheviper --number 2
Your favorite color is: blue
The magic number is: 2
```

MAGIC! 🎩✨
