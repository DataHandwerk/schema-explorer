# Sql Data Viewer

Copyright 2015-17 Tim Abell

[http://www.timwise.co.uk/sdv/](http://www.timwise.co.uk/sdv/)

An html based viewer of SQL Server Databases written in
[Go](https://golang.org/)

Supports Sqlite an MSSQL.

Note there is no protection against:

* sql injection
* cross-site-script injection (xss)

So don't give anyone access to this that you don't want to have full access to
your database. It is advised that you create a read-only database account to use with sdv.

Start the program by calling it from a shell with the path to a sqlite database:

Usage:

`./sdv-linux-x64 driver "connection" port listenAddress`

* ./sdv-linux-x64 - the supplied executable (this is the linux one, windows version also available)
* driver
  * mssql or sqlite
* connection
  * mssql - connection string
  * sqlite - path to db file
* port - optional, defaults to 8080
* listenAddress - optional, defaults to localhost which means you can onlyl see the site from the machine sdv is running on, use 0.0.0.0 to accept requests from other machines (check your firewall allows it too!)

E.g.

`./sdv-linux-x64 mssql "server=sdv-adventureworks.database.windows.net;user id=sdvRO;password=Startups 4 the rest of us;database=AdventureWorksLT" 80 0.0.0.0`

Download an example sqlite db from http://chinookdatabase.codeplex.com/ -
extract `Chinook_Sqlite_AutoIncrementPKs.sqlite` from the zip and point sdv at
it. Ignore all the build and sql files, you don't need them.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
