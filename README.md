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

# Install Atlas
```
curl -sSf https://atlasgo.sh | sh
```

# Build 
```
task
```

# Start postgres
```
docker run -p5432:5432 --name drift-db -e POSTGRES_PASSWORD=postgres -d postgres
```

# Start keydb (redis)
```
docker run -p6379:6379 --name drift-cache -d eqalpha/keydb
```

# Install Helm
```
curl https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 | bash
```

# Install k3d
```
curl -s https://raw.githubusercontent.com/k3d-io/k3d/main/install.sh | bash
```


# Setup cluster
```
k3d cluster create drift --registry-create drift-registry:0.0.0.0:5000
```
