### Liskov Substitution Principle

- Các class con có thể thay thế class cha, Nhưng không làm thay đổi đi tính đúng đăn của chương trình

Ví dụ:
- Ta có 1 class cha sử dụng ở nhiều nơi
- Nếu ta tạo ra 1 class con thừa kế từ class cha, và override methods của class cha 
- Và class con phải có thể thay thế được cho class cha,
- Nếu class con khi thay vào chỗ class cha được dùng mà không chạy được thì ta đã làm sai nguyên tắc này.
