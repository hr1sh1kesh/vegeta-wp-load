# vegeta-wp-load
Load test Wordpress using Vegeta

Deploy Wordpress with the Basic-Auth plugin first. (You can refer to the git repository https://github.com/hr1sh1kesh/wordpress-basic-auth) for the docker image for it. 

To load test the application via vegeta apply the kubernetes job in `wp-attack.yaml`
Below are the parameters that need to be passed as arguments to the Pod. The pod runs a golang cli app 

Usage: 
- `-a` the context root url of the wordpress API. (eg: `http://<nodeIP>:<nodeport>`)
- `-n` the number of requests to be generated per second.
- `-d` the duration of the test in seconds.
- `-u` the user which was created while installing wordpress via the gui. 
- `-p` the password of the user.

Example: 

`vegeta-wp-load loadgen -a http://70.0.149.108:30303/ -n 5 -d 50 -u loadgenuser -p loadgenpass`