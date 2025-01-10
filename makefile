install:
	sudo cp ./zzv /usr/local/bin/zzv
	cp ./zzv.config.json ~/zzv.config.json

uninstall:
	sudo rm /usr/local/bin/zzv
	rm ~/zzv.config.json