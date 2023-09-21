echo Start update libgosys
go build -buildmode=c-shared -o ./lib/libgosys.so ./src/system.go
mv ./lib/libgosys.h ./include/
echo Update Done