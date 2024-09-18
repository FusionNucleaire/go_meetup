# Meetup

Ce dépôt accompagne le meetup "Rust & Go, les petits nouveaux", données par Florent Le Bail et moi-même le 17/09/2024.

Les slides sont disponibles ici : https://gamma.app/docs/Rust-GO-r8ka175hkx08b03?mode=doc

Ce dépôt concerne la partie Go. Pour la partie Rust c'est par ici : https://github.com/florent-lb/rust-comet




## Comment installer Go ?

https://go.dev/doc/install => choisir son système et la bonne architecture (i686, ARM, etc...)

Vérifier avec 'go version' que Go est installé sur votre système




## Comment jouer les exemples du dépôt ?

Ouvrez un terminal (Commande Dos, Terminal Linux ou Windows, Terminal intégré à VSCode, Powershell...)

En prenant pour exemple le fichier 0 - Hello world.go : 

Pour créer un exécutable : 

**go build 0 - Hello world.go**

Cela génèrera un .exe que vous pourrez jouer par la suite

**0 - Hello world.exe**


Pour exécuter en direct le fichier sans créer d'exécutable

**go run 0 - Hello world.go**

Même chose pour les fichiers de 0 à 6

Pour le Greeter, exécuter le fichier et ouvrir la page http://localhost:10037/hello

Vous devriez avoir un Hello! dans la page en retour

Pour le main.go dans le dossier structs, se rendre dans le terminal 

**cd structs**

et faire 

**go run main.go**
