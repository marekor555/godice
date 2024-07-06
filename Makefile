default: linux windows android

linux:
	fyne package --target linux --name godice --release

windows:
	CC=x86_64-w64-mingw32-gcc GOOS=windows CGO_ENABLED=1 go build .

android:
	fyne release --target android --name godice --id com.example.godice -appVersion 1.0 -appBuild 1 -keyStore keystore.jks --keyName release-key
	fyne package --target android --name godice --id com.example.godice --release

keystore:
	keytool -genkeypair -v -keystore keystore.jks -keyalg RSA -keysize 2048 -validity 10000 -alias release-key


clean:
	rm godice.aab godice.exe godice.tar.xz godice.apk