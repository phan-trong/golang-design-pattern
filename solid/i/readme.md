### Interface Segregation Principle

- Thay vì sử dụng 1 interface lớn, thì ta nên tách ra thành nhiều interface nhỏ với những methods trong interface đó liên quan đến nhau thành những cụm để dễ dàng quản lý hơn.
- Tránh việc implement những methods không cần sử dụng vào class.

Ví dụ: 
- 1 interfaces có 100 methods bao gồm những methods như đi, đứng, chạy, nhảy, bay, bơi,
- Thì ta nên tách ra thành nhiều interface nhỏ hơn để những class như Dog, Cat, Human có thể implement những methods cần thiết của class đó tránh dư thừa những methods không cần thiết.

