# ryg
RateYourGame is an API inspired by RateYourMusic but for games. It is being developed with Go 

# Running Server
If you want to run the server localy, first you need to install Go.
Please refer to `https://go.dev/doc/install` to install Go

Second you need to install PostgreSQL to your computer. 
On Linux run 
`sudo sh -c 'echo "deb https://apt.postgresql.org/pub/repos/apt $(lsb_release -cs)-pgdg main" > /etc/apt/sources.list.d/pgdg.list'
wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | sudo apt-key add -
sudo apt-get update
sudo apt-get -y install postgresql`

After installing PostgreSQL, create a `.env` file and define `DB_NAME`, `DB_USER`, `DB_PASSWORD`.

Next, you will need to apply migrations to your database. Run `psql -U DB_USER -d DB_NAME -a -f file_name`.

This is all you need to do, now you can use RYG API to start developing your game critic site.
