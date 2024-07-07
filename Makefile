# This file is part of Godice.
# Godice is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License
# as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.
# Godice is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
# See the GNU General Public License for more details.
# You should have received a copy of the GNU General Public License along with Godice. If not, see <https://www.gnu.org/licenses/>.

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