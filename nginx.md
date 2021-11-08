```shell
nano /etc/nginx/sites-available/api-server-files.duckdns.org
```

```nginx
server {
    listen 80;
    server_name api-server-files.duckdns.org www.api-server-files.duckdns.org;

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
ln -s /etc/nginx/sites-available/api-server-files.duckdns.org /etc/nginx/sites-enabled/api-server-files.duckdns.org
nginx -t 
sudo service nginx restart 
sudo service nginx status
```

```shell
sudo certbot --nginx -d api-server-files.duckdns.org -d www.api-server-files.duckdns.org
```

