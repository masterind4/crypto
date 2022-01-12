---
marp: true
theme: default
paginate: true
footer: "Master Industrie 4.0 - 2021-2022 — Cryptographie — Guillaume Bienkowski"
math: 'mathjax'
---

<style scoped>
h1 {
  text-align: center;
}
p {
    position: absolute;
    bottom: 50px;
    right: 30px;
}
img[alt~="center"] {
  display: block;
  margin: 0 auto;
}
</style>

# TP Cryptographie 2ème partie

Guillaume Bienkowski — Braincube

---
<!-- header: "Plan du TP" -->

# TP

- Fonctions de hachage
- Signatures digitales
- Certificats

---
<!-- header: "Fonctions de hachage" -->
# Fonction de hachage

Pre

---

# Fonction de hachage parfaite $f: X \to Y$:

$\forall x \in X, \forall x' \in X$ alors $(f(x) = f(x') \implies x = x')$

On dit que $f$ est *injective*.

**Implications**: si $X$ est grand alors $Y$ sera au moins aussi vaste => ne rentre pas dans la RAM d'un ordinateur
**En pratique**: $Y$ a une taille finie (énorme, mais finie).

---

# Exemples de fonctions de hachage (imparfaites...)

- famille SHA (1, 256, ...) ($160 ≤ d ≤ 512$)
- MD5 ($d = 128$),
- Whirlpool ($d = 512$)
- BCrypt ($d=192$), Scrypt, Argon2 ($d=256$)
- $maFonction(x) = x \mod 32768$ (pas très efficace celle là)

---
# Exemple sous linux

    $ (echo "123456789" | shasum); (echo "123456779" | shasum)
                    ^____________________________^

    6d78392a5886177fe5b86e585a0b695a2bcd01a05504b3c4e38bc8eeb21e8326
    8f9d6dbc5c656b3fd63f25e72c3ec9d7738f198238a46eeb01875ee102c34860

### Petit changement en entrée => grande différence en sortie
---

# Collisions

Une collision, c'est quand deux valeurs en entrée d'une fonction de hachage donnent la même sortie.

Par exemple:

```python
>>> def mafonctiondehachage(x):
...   return x % 32768
...
>>> mafonctiondehachage(10)
10
>>> mafonctiondehachage(32778)
10 # :-( je l'avais dit qu'elle n'était pas terrible cette fonction
>>>
```

---

# Propriétés souhaitées d'une fonction de Hachage
- $f(x) = y$ est rapide à calculer
- Trouver $x$ à partir de $y$ est quasi impossible ("fonction à sens unique")
- Il doit être quasi impossible de trouver deux $x$ et $x'$ qui donnent le même *hash* (collision)
- Les valeurs retournées sont réparties uniformément dans l'espace de sortie $Y$

---
<!-- header: "Applications aux fonctions de Hachage" -->

# Applications des hash

## HashMaps, Dictionnaires (Programmation)

![](hash_table.png)

**Avantages**: $f(x)$ TRÈS rapide vs itération systématique du tableau
<!-- Faire un dessin au tableau pour expliquer l'intérêt -->
---

## Vérification d'intégrité

*Rappel*: différences minimes => hash complètement différents

![](debiandl.png)

Permet de s'assurer que le fichier (gros) qu'on a téléchargé est bien le même que celui hébergé par le serveur.

```
c685b85cf9f248633ba3cd2b9f9e781fa03225587e0c332aef2063f6877a1f0622f56d44cf0690087b0ca36883147ecb5593e3da6f965968402cdbdf12f6dd74  debian-11.2.0-amd64-netinst.iso
f2da0996196f19585b464e48bfb6b8e4938eb596667d92a5ebd428e1a88a1a115c00f1d052f350eca44fa08f42f0500c63351763dfb47f1e1783f917590c175d  debian-edu-11.2.0-amd64-netinst.iso
9b5b0475fbb3235ebb7da71415f10921b02131b7debc9325403f85f9f6798a3e902d6257831a7ec9c7aef3256817fb76c4f01bb5d035bfcdb3dc24da24fa1bda  debian-mac-11.2.0-amd64-netinst.iso
```

---

## Stockage obfusqué d'un password

On ne stocke jamais un mot de passe en clair dans une base de données:

![](non.png)
=> *NON*

On stocke un hash:

![](mouais.png)

=> *Mieux*

Si la BDD est piratée, impossible de dériver (facilement) le mot de passe.

---

### Vérification du mot de passe

![](pwd_check.png)

L'utilisateur envoie son login et mot de passe, on le hash, et on vérifie que ce hash est le bon pour cet utilisateur.

---

Attaques: **Rainbow Tables** et autres attaques basées sur un dictionnaire

Exercice mental:

- longueur password = 8
- symboles alphanumériques (`0-9`, `a-z`, `A-Z`) => 62 symboles différents
- alors on a $62^8 = 218340105584896$ combinaisons possibles
- chaque sha256 prend.. 256bits (32 octets)
- Le stockage nécessaire pour stocker un dictionnaire inverse:
  $62^8 * 32bits =6,2\nobreakspace Pio$

À la portée de n'importe quelle Fortune 500, voire du commun des mortels pour quelques heures en louant du disque chez Amazon.

Solution: *saler* son hash.
<!-- En pratique, on peut aller jusque 14 caractères sur du sha256 maintenant -->

---

### Salage

On ajoute un petit peut d'aléatoire dans le hash pour tempérer les attaques par dictionnaires.

```python
import hashlib

def sha256(s):
  return hashlib.sha256(s.encode('utf-8')).hexdigest()

password = '4gNnLar5'
salt = 'UneStringRandomDeLongueurEntre5Et16'

standard_hash = sha256(password)
# c36b9fc5e51d59f5179e9cc2a0e1f02ea6c2f12448e9e1dfe01f68786092a924

salted_hash = salt + '!' + sha256( salt + password )
# UneStringRandomDeLongueurEntre5Et16!93d4b242b475a73d8f2d1de1...
```

=> password de longueur 8, mais sha256 basé sur une longueur aléatoire.

---

## Autres applications

- [HaveIBeenPwnd](https://haveibeenpwned.com/)
- [Private Contact Discovery](https://signal.org/blog/private-contact-discovery/) (Signal)
- Blockchain
- Digital Signature (ça tombe bien c'est la suite, mais d'abord...)



----

# À vous!

Rendez-vous ici: https://github.com/masterind4/crypto

Ou directement: `git clone https://github.com/masterind4/crypto.git`

Ouvrez le fichier `hash.py` dans un éditeur de code.

---

# Signatures digitales

Revenons au téléchargement de notre fichier iso Debian:

Si ${sha256}(fichier) = S_{Controle}$ alors je sais que le fichier est *intègre* (les octets sont les mêmes que sur le serveur)

Comment vérifier l'*authenticité* du fichier? (que c'est bien Debian qui me l'a fourni)
=> La signature digitale.

![bg right fit](pirate.png)

---

## Principe
- on crypte *l'empreinte* avec une clé privée (c'est la *signature*)
- on diffuse la clé publique associée
- les gens sont capables de décrypter la signature, récupérer l'empreinte, et vérifier par eux mêmes

![bg right fit](sign.png)

---

## Vérification

On a un ISO et un fichier de signature

- on calcule le hash de l'iso
- on decrypte la signature et on obtient le hash théorique
- Si les deux sont identiques: c'est bien Debian qui a fabriqué cet ISO

![bg right fit](verify.png)

---

# Récapitulatif

- Un *hash* (ou empreinte) assure l'intégrité de la donnée
- Une signature assure *l'authenticité* de la donnée
- (L'encryption assure la *confidentialité* de la donnée, cf ce matin)

---

# Certificats


----

NOTEs:

LE tp sera en plusieurs parties:

- petit départ pour familiariser les gens avec l'exe sha256
- ensuite on passe à un D/L d'exe et on vérifie son md5sum
- on execute, ça marche.
- on D/L un autre exe, autre nom mais même md5! On execute: bam
- Fin du TP: on crée et on décortique un certificat
  - openssl
