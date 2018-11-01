Titre: Secret message sent in a multiplayer video game

Type: Forensics

Difficulté: Moyen

Durée: 45 minutes

Description: 

Deux terroristes sud américains s'envoient des message dans une biliothèque en utilisant un jeu 3D possédant un mode multijoueur ayant une fonction "chat". On sait que le jeu fonctionne par pair à pair (P2P) et utilise le protocol IPX qui doit mis dans un tunnel afin de transmettre les paquets sur le réseau.

Les transmissions entre les deux ordinateurs utilisés ont été capturées par un agent de la CIA.

Trouvez la phrase diabolique utilisée par les terroristes avant de lancer leur opération.

Format du flag: DCI{mechant message envoye}


English:

Two South American terrorists are in a library, sending messages to each other using a 3D game with a multiplayer mode which has an ingame chat. We know that the game works via peer-to-peer (P2P) and uses the IPX protocol which must be tunneled in order to transmit the network packets overt the library's network.

The data has been intercepted by a CIA agent.

Find the devilish phrase used by the terrorists before launching their operation.

Flag format: DCI{mechant message envoye}



*** DEUXIÈME DEFI -- SECOND CHALLENGE ***

Il existe un deuxième flag. Celui-ci est plus simple à trouver et prend moins de 20 minutes.

Trouvez le nom du système d'exploitation utilisé sur les ordinateurs de la bibliothèque.

Le flag est le SHA1 de son nom. Donc si le SE est DragonFly BSD, on fait le hash du string "dragonfly bsd". 
Le résultat serait: DCI{9deea0b8322b3d006494e1ab160894a1f9311c3d}


There's a second flag. It's easier to find and should take less than 20 minutes.

Find the name of the operating system used in the library. The flag is the SHA1 of its name.

For example, the if the OS was DragonFly BSD, the string to hash would be "dragonfly bsd". 
The resulting flag would be: DCI{9deea0b8322b3d006494e1ab160894a1f9311c3d}



