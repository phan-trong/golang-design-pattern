### Open-Closed Principle

- Không nên chỉnh sửa class cha, nên tạo class con kế thừa class cha và mở rộng nó.
- Vì khi sửa class cũ có khả năng ảnh hưởng đến những code cũ. miss những func ở những nơi mình sử dụng class đó.

Ví dụ:
- 1 hàm trong class cha trả vễ chuỗi, và sử dụng hàm đó ở khoảng 50 nơi khác
=> khi này nếu ta sửa class cha trả về int thì 50 nơi sử dụng hàm đó sẽ hỏng hết không sử dụng được nữa.
