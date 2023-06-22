# TechnicalTest Golang K-Style Hub

Name : Aria shabry (aria.shabry@gmail.com)



**Setting Environtment**

1. Update your environtment at env.yaml file like DB_Host, DB_PORT, DB_NAME, and etc.
2. If your environtment error, please check env.yaml file and ConnectionString at ./helpers/env/env.go
3. Database Name = skincare
4. Import Database using skincare.sql file
5. Application address : http://localhost:5000/api/

**Answer**


1. Insert data ke tabel Member
- screenshot hasil : ./1-CreateMember.png
![CreateMember](https://github.com/Ariashabry/TechnicalTest/blob/main/1-CreateMember.png?raw=true)
- api : POST/http://localhost:5000/api/member
- request body : raw/json
- example :
```
{
    "Username": "Puan Maharani",
    "Gender": "perempuan",
    "SkinType": "normal",
    "SkinColor": "light" 
}
```

2. Update data tabel Member
- screenshot hasil : ./2-UpdateMember.png
![UpdateMember](https://github.com/Ariashabry/TechnicalTest/blob/main/2-UpdateMember.png?raw=true)
- api : PUT/http://localhost:5000/api/member
- request body : raw/json
- example
```
{
    "Username": "Puan Maharani",
    "Gender": "perempuan",
    "SkinType": "oily",
    "SkinColor": "light" 
}
```

3. Delete data tabel Member
- screenshot hasil : ./3-DeleteMember.png
![DeleteMember](https://github.com/Ariashabry/TechnicalTest/blob/main/3-DeleteMember.png?raw=true)
- api : DELETE/http://localhost:5000/api/member/:id
- :id merupakan id_member

4. Select all tabel Member
- screenshot hasil : ./4-GetAllMember.png
![GetAllMember](https://github.com/Ariashabry/TechnicalTest/blob/main/4-GetAllMember.png?raw=true)
- api : GET/http://localhost:5000/api/member

5. Select product by ID_PRODUCT
- screenshot hasil : ./5-GetProductById.png
![SelectProductById](https://github.com/Ariashabry/TechnicalTest/blob/main/5-GetProductById.png?raw=true)
- api : GET/http://localhost:5000/api/product/:id
- :id merupakan id produk

6. Insert Like Review
- screenshot hasil : ./6-InsertLikeReview.png
![Like](https://github.com/Ariashabry/TechnicalTest/blob/main/6-InsertLikeReview.png?raw=true)
- api : POST/http://localhost:5000/api/like
- request body : raw/json
- example
```
{
    "IdReview": 2,
    "IdMember": 1 
}
```

7. Delete Cancel Like Review
- screenshot hasil : ./7-DeleteCancelLikeReview.png
![Unlike](https://github.com/Ariashabry/TechnicalTest/blob/main/7-DeleteCancelLikeReview.png?raw=true)
- api : DELETE/http://localhost:5000/api/cancel-like
- request body : raw/json
- example
```
{
    "IdReview": 2,
    "IdMember": 1 
}
```

Jika ada pertanyaan, silahkan hubungi saya di : aria.shabry@gmail.com


