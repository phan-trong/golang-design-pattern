# Design Pattern Abstract Factory

## Mục đích sử dụng

- Dùng để tạo ra một đối tượng ở trong 1 họ các đối tượng liên quan đến nhau mà không cần chỉ định đối tượng cụ thể
- Hiểu nôm na là mô hình này to hơn factory methods, ở trong abstract method chưa nhiều factory methods.

## Vấn đề
Ví dụ:
- Ta có 1 họ bàn ghế gồm: Bàn Coffee, Ghế đơn, Ghế Dài.
- Có nhiều họ tương tự: Cổ điển, Hiện Đại, Victorian.
![image description solution](./problem-en.png)

## Giải Pháp

- Đầu tiên tạo ra interface của mỗi sản phẩm của mỗi họ riêng biệt,
![image description solution](./solution1.png)
- Tiếp theo định nghĩa Abstract Factory, một interface có những methods để tạo ra tất cả các sản phẩm của họ đó.
![image description solution](./solution2.png)

## Cách dùng:
- Làm matrix tất cả sản phẩm riêng biệt so với các biến thể của các sản phẩm này.
- Sau đó tạo ra interface của tất cả sản phẩm, và tạo ra tất cả các concrete class và implement tất cả những methods đó.
- Tạo ra factory class, để tạo đối tượng cho mỗi họ (Họ: là những class liên quan đến nhau)
- Tạo mội class để khởi tạo các factory

### Nguồn: 
https://refactoring.guru/design-patterns/abstract-factory
https://refactoring.guru/design-patterns/factory-method/go/example