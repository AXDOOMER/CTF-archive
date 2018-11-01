# Linux.Trojan (500 pts)

Un nouveau logiciel malveillant a été découvert par une firme en Europe de l'Est. Il est droppé par un script Bash d'un paquet communautaire de Open Suse modifié par un script kiddie et ayant passé inaperçu. Les détections sont multiples; les spécialistes craingnent que ce soit un nouveau botnet et que sont réseau compte déjà plus de 36000 bots. C'est presque 9% des utilisateurs de Open Suse qui sont infectés.

Son nom jusqu'à présent est Linux.Trojan, mais personne ne l'a encore rétro-ingénié pour savoir exactement ce qu'il fait et lui donner une catégorisation plus spécifique. Il s'agit d'un prototype qui pourrait être bientôt mis à jour.

Il est à noter que ce faux maliciel utilise des techniques similaires à celles utilisées par des vrais maliciels afin de rendre votre tâche d'analyse plus complexe. Au moins les symboles de débogage ait été mis à l'intérieur pour vous aider.

Le flag suit le format DCI{ }.

# Solution

Il s'agit d'un faux maliciel avec plusieurs fonctionalités qui empêchent son analyse. Il dispose d'une protection anti-VM rudimentaire, ainsi qu'un moyen de détecter des outils d'analyse (Wireshark et tcpdump).

La solution la plus simple est d'enlever la protection anti-VM et de rouler Wireshark à l'extérieur de sa machine virutuelle. Celui-ci pourra capturer le traffic réseau et le maliciel ne pourra pas le détecter. Le code de la protection anti-VM n'est pas obfusqué et est caché dans la fonction "datetime".

Il y a une fonction qui sert à l'encryption. L'algorithme utilisé est TEA. Avec un débogueur, il est possible de récupérer la clef d'encryption et voir ce qui est décrypté. Il s'agit d'URLs et de mot clefs qui déclenchent un contact avec le serveur C2. Le flag est retourné par le C2 et sera intercepté dans Wireshark lorsqu'un de ces mots clefs est entré au clavier.

Flag: DCI{Linux.Trojan.Not.A.Keylogger}

*Aucune personne n'a réussi ce challenge.*

# Solution 2

Si on a le matériel, on peut se faire un setup baremetal pour tester le maliciel. Un analyste de maliciels suivant la procedure d'analyse dynamique pourrait résoudre le défi en moins de 30 minutes.

Il s'agit d'avoir un ordinateur et de connecter un câble réseau à un autre qui sert de proxy. Celui-ci peut intercepter le traffic réseau et le faux maliciel n'en aurait pas connaissance. Ce deuxième ordinateur serait connecté à Internet et la connexion serait forwardée jusqu'au premier ordinateur. Wireshark pourrait capturer tout traffic réseau passant à travers le proxy à ce moment là.

Quelqu'un qui connait l'existence des credential stealers et banking trojans (souvent des trucs retrouvés chez des maliciels de type botnet genre TrickBot) essayerait à ce moment des sites comme Facebook et des sites de banques afin d'examiner le comportement du maliciel. Ce faux maliciel attend qu'on entre un nom de banque canadienne ou le nom d'un autre site intéressant avant de contacter le server C2. Le serveur C2 retourne le flag à ce moment là.
