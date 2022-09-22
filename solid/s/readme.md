### Single Responsibility Principle

- Mỗi class chỉ nên làm một nhiệm vụ duy nhất và chỉ có một lý do để sửa class đó
- Nên tách nhỏ những class ra để dễ maintain cho sau này, dễ đọc hơn.

Ví dụ:
- Ở mô hình MVC
+ Model chỉ có nhiệm vụ tương tác database
+ View chỉ có nhiệm vụ hiển thị dữ liệu cho người dùng
+ Controller chỉ có nhiệm vụ nhận request, xử lý dữ liệu lầy từ Model và trả về

=> Nếu Model có xử lý logic xử lý dữ liệu ở tầng Model, thì đã vi phạm nguyên tắc này.
