xgo --targets=windows/amd64 --trimpath --ldflags \
  "-s -w -X main.version=0.6.0 -X main.commitID=$(git rev-parse --short HEAD 2>/dev/null)" github.com/aaron-skillz/sync-server-go
