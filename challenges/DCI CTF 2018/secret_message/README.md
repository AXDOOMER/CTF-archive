# secret_message_part1 (50 pts) et secret_message_part2 (200 pts)

Deux terroristes sud américains s'envoient des message dans une biliothèque en utilisant un jeu 3D possédant un mode multijoueur ayant une fonction "chat". On sait que le jeu fonctionne par pair à pair (P2P) et utilise le protocol IPX qui doit mis dans un tunnel afin de transmettre les paquets sur le réseau.

Les transmissions entre les deux ordinateurs utilisés ont été capturées par un agent de la CIA.

flag part 1: Trouvez le nom du système d'exploitation utilisé sur les ordinateurs de la bibliothèque. Le flag est le SHA1 de son nom. Donc si le SE est DragonFly BSD, on fait le hash du string "dragonfly bsd". Le résultat serait: DCI{9deea0b8322b3d006494e1ab160894a1f9311c3d}

flag part2: Trouvez la phrase diabolique utilisée par les terroristes avant de lancer leur opération. Si la phrase était: "J'adore les pêches", le flag serait: DCI{j'adore les peches}

# Solution

### Partie 1

Pour le premier flag, qui est le plus facile à trouver, il faut examiner les requêtes DHCP. On peut trouver le hostname "voidli" dans le paquet No. 12650. Si on ne connaît pas cette distribution Linux, une recherche sur Google nous revèle qu'il s'agit de Void Linux. On fait donc le hash de "void linux" pour obtenir 

DCI{aa598b8fe89f5e6069f1b9019b0f869eb23e1d10}

*Il a été réussi par une personne.*

### Partie 2

Après avoir lu la description du challenge, l'information révélée nous permet de débuter une recherche sur Google. Si on cherche "3d game multiplayer ipx" qui sont des mots clefs venant de la description du challenge, le premier résultat nous amène sur [https://www.dedoimedo.com/games/reviving/dosbox_multiplayer.html](https://www.dedoimedo.com/games/reviving/dosbox_multiplayer.html). On nous montre Doom, qui correspond à la description donnée du jeu. Des instructions sont aussi fournies afin de lancer le jeu dans DOSBox (dont les match par réseau fonctionnent par IPX).

On peut télécharger la version Shareware de Doom et le tester en multijoueur dans DOSBox en suivant les instructions de ce site. En chattant et en capturant le traffic réseau en même temps, on peut voir comment le chat est trasmit. Ceci nous permet de savoir à quelle position de chaque paquet regarder dans le PCAP qui nous a été fournit. Il s'agit de caractères qui sont transmits un par un à mesure qu'ils sont tappés au clavier.

Le message commence au paquet No. 6777 ("L"), 6803 ("A"), 6807 ("S"), 6875 (" ") ... et le dernier caractère se trouve dans le paquet No. 8765 ("S").

DCI{las zanahorias estan cocidas}

*Aucune personne n'a réussi ce challenge.*
