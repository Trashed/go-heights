APP_NAME="go-heights";
# Mac (darwin) build
mkdir -p build/mac;
env GOOS=darwin GOARCH=amd64;
go build;
mv $APP_NAME build/mac/;
# Linux build
mkdir -p build/linux64;
env GOOS=linux GOARCH=amd64;
go build;
cp $APP_NAME build/linux64/$APP_NAME;
# Windows build
mkdir -p build/win64;
env GOOS=windows GOARCH=amd64;
go build -o $APP_NAME.exe;
mv $APP_NAME.exe build/win64/;
