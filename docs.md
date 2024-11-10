## 1. install go 
```bash
sudo apt update
```

```bash
wget https://go.dev/dl/go1.20.7.linux-amd64.tar.gz
```

```bash
sudo tar -xzf go1.20.7.linux-amd64.tar.gz
```

```bash
sudo mv go/bin/* /usr/local/bin
```

```bash
go version
```

```bash
sudo rm -rf go
sudo rm go1.20.7.linux-amd64.tar.gz
```

```bash
go mod init main
```

```bash
go run .
```

