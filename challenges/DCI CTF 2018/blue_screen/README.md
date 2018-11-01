# blue_screen (250 pts)

Lien de téléchargement de la dump: [dump.zip](https://drive.google.com/file/d/1Qi4hDRy4tJnIZrNF_lvwHdxcRZ_oCNU3/view?usp=sharing)

# Description

Timmy travaillait sur un fond d'écran quand son ordi a planté avec un écran bleu disant: driver irql_less_or_not_equal
L'ordi a fait une vidange de mémoire sur le disque. Est-ce que tu peux aider Timmy à le récupérer?

Timmy was working on a new wallpaper when his computer crashed with a blue screen saying: driver irql_less_or_not_equal
The computer made a physical memory dump to the disk. Can you help Timmy recover his wallpaper?

# Solution

Utiliser Volatility et extraire le processus correspondant à "Paint". On peut utiliser GIMP pour lire la memory dump du processus en tant qu'image brute. On pourra trouver une image de Donald Trump avec le flag inscrit dessus. 

Cependant, une personne a rapporté qu'il est possible de faire un `grep` sur la dump du processus et d'obtenir le flag.

flag{ronalddrupmmffff}

*Ce challenge a été réussi par 4 personnes.*
