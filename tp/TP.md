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
- Certificats et anatomie d'une connection TLS

---
# Récupération des sources

```
git clone https://github.com/masterind4/crypto.git
code crypto/
```

---
<!-- header: "Fonctions de hachage" -->

# Fonctions de hachage

## Exercice 1

Rendez-vous sur https://prometheus.io/download/

Téléchargez la version linux `amd64` de `pushgateway`

Vérifiez à l'aide de l'utilitaire `shasum` que la somme de contrôle est bonne

Extrayez et lancez en ligne de commande l'utilitaire `pushgateway --help`

----


# Fonctions de hachage

## Exercice 2

Rendez-vous sur https://masterind4.github.io/

Téléchargez la version 1.0 de `Antivirus.exe`
**UTILISEZ LE MIROIR 1**

Vérifiez à l'aide de l'utilitaire `md5sum` que la somme de contrôle est bonne

Lancez `Antivirus.exe`

----

# Fonctions de hachage

## Exercice 2 bis

Rendez-vous sur https://masterind4.github.io/

Téléchargez la version **MIROIR 2** de `Antivirus.exe`

Vérifiez à l'aide de l'utilitaire `md5sum` que la somme de contrôle est bonne

Lancez `Antivirus.exe`

Que s'est-il passé?

---

<!-- header: "Signatures digitales" -->

TODO: vérifier la signature d'un paquet en utilisant gpg

---

<!-- header: "Certificats" -->


# Certificats

## Exercice 1: Utiliser `openssl s_client` pour se connecter de façon sécurisée à un serveur tiers.

Tenter une connection sur `neolyse.info` sur le port `443`, et utiliser l'option `-showcerts` pour afficher les certificats renvoyés par le serveur.

Vérifier que la connection retourne bien `Verify return code: 0 (ok)`, qui signifie que la chaîne de certification est bien valide.

Aide:
```man s_client```

---

Sauvegarder chaque certificat affichés par la commande précédente dans un fichier avec l'extension `.pem` et faire afficher avec `openssl x509` les détails de chaque certificats (utiliser l'option `-text` pour afficher les détails en format lisible.)

Récupérer la liste des DNS autorisés (`Subject Alternative Name`) par le certificat final.

Constater les contraintes (`Basic Constraints`) sur les certificats intermédiaires, et sur le certificat final (`CA:FALSE`)
