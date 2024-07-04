default: linux windows android

linux:
	fyne package --target linux --name godice --release

windows:
	CC=x86_64-w64-mingw32-gcc GOOS=windows CGO_ENABLED=1 go build .

android:
	fyne package --target android --name godice --id com.example.godice --release