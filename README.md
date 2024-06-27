### Uyga Vazifa: Unary gRPC Interceptorlarini Go tilida amalga oshirish

#### Maqsad
Ushbu vazifaning maqsadi `gRPC` cancellation funksiyasini vazifa boshqaruv tizimida tushunish va tatbiq etishdir. Vazifa boshqaruv tizimi foydalanuvchilarga yangi vazifalar yaratish, ro'yxatga olish va cancellation imkonini beradi.

#### Vazifalar

1. **gRPC Serviceni yaratish**:
   - Hech bo'lmaganda bitta usulga ega bo'lgan oddiy `gRPC` serviceni amalga oshiring, masalan `GetProduct`.
   - Xizmat uchun protobuf faylini yarating
   
   Misol uchun `.proto` file:
   ```proto
    syntax = "proto3";

    option go_package = "./productpb";

    service ProductService {
        rpc GetProduct (ProductRequest) returns (ProductResponse);
    }

    message ProductRequest {
        string id = 1;
    }

    message ProductResponse {
        string id = 1;
        string name = 2;
        double price = 3;
    }
    ```

2. **Serverni amalga oshirish:**:

3. **Unary Interceptorlarini amalga oshirish:**:
   - `Logging Interceptor`: Har bir kelgan so'rov va chiqayotgan javob tafsilotlarini log qiladigan unary interceptor yarating.
  

4. **Interceptorni integratsiya qilish:**:
  - Logging interceptorini `gRPC` serveringizga integratsiya qiling.
  
5. **Clientni amalga oshirish:**:
  - `gRPC` servicega so'rovlar yuboradigan va javoblarni qayta ishlaydigan clientni amalga oshiring

6. **Test qiling**:
  - Client tomonidan to'g'ri va noto'g'ri so'rovlarni yuborib, `gRPC` serviceni va interceptorni sinab ko'ring.
  - Logging interceptor kerakli ma'lumotlarni log qilishini tekshiring.

7. **Qo'shimcha Vazifa**:
  - `AddProduct`, `DeleteProduct` `gRPC` servicelarni amalga oshiring


### O'qim
```
+-----------------------------------------+
|           Client Application            |
+-----------------------------------------+
            |
            | 1. Make a request
            v
+-----------------------------------------+
|          gRPC Client Library            |
+-----------------------------------------+
            |
            | 2. Pass request along
            v
+-----------------------------------------+
|        Logging Interceptor              |
|  - Logs the details of the request      |
+-----------------------------------------+
            |
            | 3. Pass request along
            v
+-----------------------------------------+
|           gRPC Server Library           |
+-----------------------------------------+
            |
            | 4. Route request to method
            v
+-----------------------------------------+
|         ProductService Server           |
|  - Handles GetProduct request           |
|  - Returns ProductResponse              |
+-----------------------------------------+
            |
            | 5. Pass response back
            v
+-----------------------------------------+
|        Logging Interceptor              |
|  - Logs the details of the response     |
+-----------------------------------------+
            |
            | 6. Pass response back
            v
+-----------------------------------------+
|          gRPC Client Library            |
+-----------------------------------------+
            |
            | 7. Return response
            v
+-----------------------------------------+
|        Client Application               |
+-----------------------------------------+
```

