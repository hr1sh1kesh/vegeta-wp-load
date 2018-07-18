# vegeta-wp-load
Load test Wordpress using Vegeta

Deploy Wordpress with the Basic-Auth plugin first. (You can refer to the git repository https://github.com/hr1sh1kesh/wordpress-basic-auth) for the docker image for it. 

Execute the CLI app using

wp-attack loadgen -a <Base API path> -u <username> -p <password> -n <rate of requests/sec> -d <duration of test>
