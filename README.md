# Sql Data Viewer

Copyright 2015 Tim Abell

[http://www.timwise.co.uk/sdv/](http://www.timwise.co.uk/sdv/)

An html based viewer of SQL Server Databases written in
[Go](https://golang.org/)

Currently supports sqlite, but planned to support sql server.

Support for mysql, sql server, oracle, postgres etc will depend on demand.

Note there is no protection against:

* sql injection
* cross-site-script injection (xss)

So don't give anyone access to this that you don't want to have full access to
your database.

Start the program by calling it from a shell with the path to a sqlite database:

    ./sdv some.db

Download an example sqlite db from http://chinookdatabase.codeplex.com/ -
extract `Chinook_Sqlite_AutoIncrementPKs.sqlite` from the zip and point sdv at
it. Ignore all the build and sql files, you don't need them.