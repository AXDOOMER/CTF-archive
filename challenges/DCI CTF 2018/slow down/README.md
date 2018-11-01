# Slow down (150 pts)

Format du flag: flag{...}

# Solution

Ouvrir l'ELF dans Radare et NOPer les sleep ou modifier le paramètre de ces fonctions pour qu'il soit '0'.

Il y a deux "sleep" dans la boucle principale et un "usleep" caché dans `cout_`. 

Sur un ordi lent, ça peut prendre jusqu'à 10 secondes pour que le flag soit affiché.

Flag: flag{Daedalus+Icarus=Helios}

*Ce challenge a été réussi par 4 personnes.*
