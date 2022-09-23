# Design Pattern Factory Method

## Mục đích sử dụng

- Là một cách có thể tạo đối tượng từ nhiều class con thừa kế chung 1 class cha nhưng chưa biết cụ thể là sẽ tạo class con nào.


## Vấn đề
Ví dụ:
- Thử tưởng tượng có một hệ thống vấn tải đường bộ, bằng xe Trucks. Sau một thời gian, có nhiều yêu cầu vận chuyển bằng đường thủy.
Ta có thể thêm lớp Ships vào nhưng việc này sẽ làm thay đổi toàn bộ codebase hiện tại, Và nếu sau này có thêm vận chuyển đường hàng không
thì có thể ta sẽ phải làm lại những bước thay đổi này

## Giải Pháp

- Ta có 1 class factory và việc xây dựng đôi tượng trả về sẽ nằm ở trong class này
![image description solution](./solution1.png)
- Các sub class nên có chung 1 lớp interface
![image description solution](./solution2.png)
- ví dụ như ở trên class Truck và class Ship nên implement từ interface Transport. Khai báo 1 methods deliver. Môi class sẽ implement phương thức này một cách khác nhau.
![image description solution](./solution3.png)

## Cách dùng:
- Làm tất cả các class có chung interface,
- tạo ra 1 class Factory có methods check type sub class muốn tạo và return về class muốn tạo.

### Nguồn: 
https://refactoring.guru/design-patterns/factory-method
https://refactoring.guru/design-patterns/factory-method/go/example