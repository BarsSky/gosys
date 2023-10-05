echo Start update libgosys
go build -buildmode=c-shared -o ./lib/libgosys.dll ./src/system.go
move ./lib/libgosys.h ./include/
echo Update Done