### Dependency Inversion Principle

- Các module cấp cao không nên phụ thuộc vào các module cấp thấp
- Module User, Cart là module cấp cao, 
- Module MysqlConnection, LocalStorage là những module cấp thấp.
- Mà nên phụ thuộc vào các interface của module cấp thấp.
- Những entities khác nhau nên phụ thuộc vào những interface khác nhau mà không phải concretions,
- Chỉ nên phụ thuộc vào interface không nên phụ thuộc vào implementation.

Ví dụ:
- Ở trong web có controller và repository. Ở trong Controller không nên sử dụng thẳng class của service mà chỉ nên có interface của service đó thôi
- Nếu sử dụng class mà không thông qua interface thì khi thay đổi class sẽ bị ảnh hưởng đến những nơi sử dụng class đó. Thay vì vậy thì
- Khi sử dụng interface nếu muốn thay đổi class, ta chỉ cần tạo ra class mới và implement những methods của interface và gọi ở nơi ta muốn.
