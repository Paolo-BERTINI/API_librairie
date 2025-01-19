# __Projet My_API__

## 1. ___Comment fonctionne le projet ?___

Le projet est composé de plusieurs dossiers:
* Le dossier `routes` sert à stocker les routes de chacunes des fonctions executants le code.
* Le dossier `models` sert à stocker les structures et variables utilisées dans le code.
* Le dossier `config` sert à configurer la connection entre le code et la base de donnée *MariaDB*.
* Le dossier `controllers` sert à stocker les fonctions permettant le fonctionnement de l'API.

gfxn
## 2. ___Comment faire fonctionner le code ?___

Pour éxécuter le code, lancer le fichier main (`go run main.go`). Ensuite se connecter aux routes souhaitées.


## 3. ___Comment utiliser l'API ?___

* ### __Routes disponibles sans être connecté:__

  * `/api/users` avec la méthode __PUT__ permet de créer son propre utilisateur dans la base de donnée en envoyant sous format JSON un prénom *first_name*, un nom de famille *last_name*, un email *mail* et un mot de passe *password* qui sera hachagé pour plus de sécurité.
  > *Exemple:*
  > ```
  > {
  >     "first_name":"Prénom",
  >     "last_name":"NOM",
  >     "mail":"mail@hotmail.com",
  >     "password":"mot_de_passe"
  > }
  > ```

  * `/api/users/connect` avec la méthode __GET__ permet de se connecter au compte utilisateur en envoyant sous format JSON son email *mail* et son mot de passe *password*.
  > *Exemple:*
  > ```
  > {
  >     "mail":"mail@gmail.com",
  >     "password":"mot_de_passe"
  > }
  > ```

  #### Voir les auteurs:

  * `/api/authors` permet avec la méthode __GET__ d'avoir le nom de tous les auteurs disponibles dans la base de donnée.
  * `/api/authors/{id}` permet avec la méthode __GET__ d'avoir le nom, la date de naissance et la description d'un auteur, spécifié par son id à la fin de la route.

  #### Voir les livres:

  * `/api/books` permet avec la méthode __GET__ d'avoir le titre et l'auteur de tous les livres encore en stock dans la base de donnée.
  * `/api/books/{id}` permet avec la méthode __GET__ d'avoir le titre, l'auteur, le résumé et le prix d'un livre, spécifié par son id à la fin de la route. Si il n'est plus en stock, *plus en stock* sera renvoyé à l'utilisateur.

<br>

* ### __Routes disponibles en étant connecté en temps qu'utilisateur:__
        
  * `/api/orders/create` permet avec la méthode __POST__ de faire un achat de livre en envoyant sous format JSON la quantité *quantity* et l'id du livre *book* souhaité. Si le livre demandé n'est plus en stock, alors *Quantité disponible insuffisante* sera renvoyé à l'utilisateur.
  > *Exemple:*
  > ```
  > {
  >     "quantity":"4",
  >     "book":"4"
  > }
  > ```

  * `/api/orders/history` permet à l'utilisateur avec la méthode __GET__ l'accès à son historique d'achat.

  #### Accès à ses données:

  * `/api/users` permet avec la méthode __PUT__ de modifier ses propres données en envoyant ses nouvelles données sous format JSON. Un prénom *first_name*, un nom de famille *last_name*, un email *mail* et un mot de passe *password* seront attendus.
  > *Exemple:*
  > ```
  > {
  >     "first_name":"Prénom",
  >     "last_name":"NOM",
  >     "mail":"mail@hotmail.com",
  >     "password":"mot_de_passe"
  > }
  > ```
  
  * `/api/users` permet avec la méthode __DELETE__ de supprimer son compte utilisateur.
  * `/api/users` permet avec la méthode __GET__ d'affiche ses propres informations.

<br>

* ### __Routes disponible en étant connecté en temps qu'administrateur:__

  #### Accès aux données des utilisateurs:

  * `/api/users` permet avec la méthode __GET__ d'afficher la liste de tous les utilisateurs présents dans la base de donnée.
  * `/api/users/{id}` permet avec la méthode __GET__ d'avoir les information d'un utilisateur, spécifié par son id à la fin de la route.
  * `/api/users/{id}` permet avec la méthode __PUT__ de modifier les informations d'un utilisateur, spécifié par son id à la fin de la route.
  * `/api/users/{id}` permet avec la méthode __DELETE__ de supprimer le compte d'un utilisateur, spécifié par son id à la fin de la route.

  #### Accès aux auteurs:

  * `/api/authors` permet avec la méthode __POST__ d'ajouter un auteur à la base de donnée en envoyant sous format JSON son nom *name*, sa date de naissance *birth_date* sous format *AAAA-MM-JJ* et une description *description*.
  > *Exemple:*
  > ```
  > {
  >     "name":"Prénom NOM",
  >     "birth_date":"AAAA-MM-JJ",
  >     "description":"description sur l'auteur"
  > }
  > ```

  * `/api/authors/{id}` permet avec la méthode __PUT_ de modifier les informations d'un auteur, spécifié par son id à la fin de la route, en envoyant sous format JSON son nom *name*, sa date de naissance *birth_date* et une description *description*
  > *Exemple:*
  > ```
  > {
  >     "name":"Prénom NOM",
  >     "birth_date":"AAAA-MM-JJ",
  >     "description":"description sur l'auteur"
  > }
  > ```

  * `/api/authors/{id}` permet avec la méthode __DELETE__ de supprimer un auteur, spécifié par son id à la fin de la route.
  **ATTENTION**: On ne peut pas supprimer un auteur dont un livre est encore présent dans la base de donnée.

  * `/api/books` permet avec la méthode __POST__ d'ajouter un livre à la base de donnée en donnant sous format JSON un titre *title*, l'id de l'auteur *author*, le stock *stock* disponible, la date de publication *publication_date* sous format *AAAA-MM-JJ*, le prix *price* et le résumé *summary* du livre.
  **ATTENTION**: L'auteur ayant écrit le livre doit être présent dans la base de donnée pour pouvoir ajouter le livre.
  > *Exemple:*
  > ```
  > {
  >     "title":"Titre",
  >     "author":"L'id de l'auteur",
  >     "stock":"Stock",
  >     "publication_date":"AAAA-MM-JJ",
  >     "price":"Prix",
  >     "summary":"Résumé"
  > }
  > ```

  * `/api/books/{id}` permet avec la méthode __PUT__ de modifier les information d'un livre, spécifié par son id à la fin de la route, en envoyant sous format JSON un titre *title*, l'id de l'auteur *author*, le stock *stock* disponible, la date de publication *publication_date*, le prix *price* et le résumé *summary* du livre.
  > *Exemple:*
  > ```
  > {
  >     "title":"Titre",
  >     "author":"L'id de l'auteur",
  >     "stock":"Stock",
  >     "publication_date":"AAAA-MM-JJ",
  >     "price":"Prix",
  >     "summary":"Résumé"
  > }
  > ```

  * `/api/books/{id}` permet avec la méthode __DELETE__ de supprimer un livre de la base de donnée, spécifié par son id à la fin de la route.
  * `/api/orders/history/{id}` permet avec la méthode __GET__ d'avoir la liste des achats de d'un utilisateur, spécifié par son id à la fin de la route.
