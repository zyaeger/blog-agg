# Gator Blog Aggregator

## Requirements & Set up

You'll need to have PostgreSQL and Go installed to run the program.
To install the `gator` CLI, you can using:
```
go install github.com/zyaeger/blog-agg@latest
```

In order to set up the config file, create a `~/.gatorconfig.json`. Make sure it lives in your root/home directory.

Inside the file, format a json with the following fields:
```
{
    "db_url": "<connection_string>",
    "current_user_name": "<username>"
}
```

That's it!