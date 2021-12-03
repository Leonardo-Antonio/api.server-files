```shell
nano /etc/nginx/sites-available/server-file.duckdns.org
```

```nginx
server {
    listen 80;
    server_name server-file.duckdns.org www.server-file.duckdns.org;

    location / {
        proxy_pass http://localhost:8001;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection 'upgrade';
        proxy_set_header Host $host;
        proxy_cache_bypass $http_upgrade;
    }
}
```

```shell
nginx -t
```




```shell
ln -s /etc/nginx/sites-available/server-file.duckdns.org /etc/nginx/sites-enabled/server-file.duckdns.org
nginx -t 
sudo service nginx restart 
sudo service nginx status
sudo systemctl enable nginx
```

```shell
sudo certbot --nginx -d server-file.duckdns.org -d www.server-file.duckdns.org
```

