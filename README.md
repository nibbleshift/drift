# drift

# Install Taskfile
```
sudo sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b /usr/local/bin
```

# Install Node

Ubuntu
```
curl -fsSL https://deb.nodesource.com/setup_20.x | sudo -E bash - &&\
sudo apt-get install -y nodejs
```

# Install yarn

```
sudo npm install --global yarn
```

# Build 
```
task
```

# Start postgres
```
docker run -p5432:5432 --name drift-postgres -e POSTGRES_PASSWORD=postgres -d postgres
```
