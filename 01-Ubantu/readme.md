# Installing HAProxy on Ubuntu

Before proceeding the installation of HAProxy, you are advised to update your system by typing the command below:
```
sudo apt-get update && sudo apt-get upgrade
```

Installation:
To install HAProxy, type the command below followed by y when prompted to continue:
```
sudo apt-get install haproxy
```

Afterwards, you can double-check the installed version number with the following command.
```
haproxy -v
```

## Configuring the load balancer
The HAProxy configuration process involves three(3) major sources of parameters :
The arguments from the command-line, which always take precedence
The “global” section, which sets process-wide parameters
The proxies sections which can take form of “defaults”, “listen”, “frontend” and “backend”.


Setting up HAProxy load balancer is a quite straightforward process. Basically, all you need to do is tell HAProxy what kind of connections it should be listening for and where the connections should be relayed to.

### Load balancing at layer 4
Start off with a basic setup. Create a new configuration file, for example, using vi with the command underneath.
```
sudo vi /etc/haproxy/haproxy.cfg
```

Add the following sections to the file. Replace the server_name with whatever you want to call your servers on the statistics page and the private_ip with the private IPs for the servers you wish to direct the web traffic. You can check the private IPs at your UpCloud control panel and Private network tab under Network menu.

```json
frontend http_front
   bind *:80
   stats uri /haproxy?stats
   default_backend http_back

backend http_back
   balance roundrobin
   server server_name1 private_ip1:80 check
   server server_name2 private_ip2:80 check
```

After making the configurations, save the file and restart HAProxy with the next command.
```
sudo systemctl restart haproxy
or
sudo service haproxy restart
```

Testing the setup
With the HAProxy configured and running, open your load balancer server’s public IP in a web browser and check that you get connected to your backend correctly. The parameter stats uri in the configuration enables the statistics page at the defined address.
```
http://load_balancer_public_ip:8443/haproxy?stats
```

# Set SSL
```
sudo openssl req -x509 -nodes -days 1365 -newkey rsa:2048 -keyout /home/nutsu/ssl_certificate/haproxy.key -out /home/nutsu/ssl_certificate/haproxy.crt
```

## Creating a Combined PEM SSL Certificate/Key File
```
sudo cat haproxy.crt haproxy.key > haproxy.pem
sudo cp haproxy.pem /home/nutsu/ssl_certificate/
```


# Configuring Rsyslog to Collect HAProxy Logs
```
sudo vi /etc/rsyslog.d/99-haproxy.conf
```

Add config below:
```
$AddUnixListenSocket /var/lib/haproxy/dev/log

# Send HAProxy messages to a dedicated logfile
:programname, startswith, "haproxy" {
  /var/log/haproxy.log
  stop
}
```

then you can restart Rsyslog with the following command:
```
sudo service rsyslog restart
```

Testing HAProxy Logging

Now that you have configured HAProxy, Rsyslog, and optionally SELinux, you can test that logging to /var/log/haproxy.log is working.
```
tail -f /var/log/haproxy.log
```