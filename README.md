#RUN REDIS in terminal
docker run -p 6379:6379 redis:latest 

#install existing pacakge
go mod tidy

#RUN REDIS
Once you’ve installed Node.js and npm, it’s a simple one-liner to get and install the Node.js version of redis-cli:

npm install -g redis-cli

Then you can run it with the command:

rdcli -h your.redis.host -a yourredispassword -p youredishost

rdcli -h 127.0.0.1 -a password -p 6379

rdcli -h 127.0.0.1 -p 6379