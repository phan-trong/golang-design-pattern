# Inversion of Control (IoC)

- Là một nguyên tác lập trình. sinh ra để tuân
- Như tên gọi của nó, thì nó được dùng đẻ đảo ngược sự phụ thuộc trong thiết kế hướng đối tượng,

## Giải thích
 Ví dụ:
- Hằng ngày bạn phải lái xe đi làm thì bạn đang điều khiển xe để đi đến nơi làm việc. 
- Thay vào đó bạn có thể thuê taxi, grab để tài xế lái xe giúp bạn, Vậy là sẽ có một người khác lái xe giúp bạn
và bạn không bị phụ thuộc vào xe của mình nữa. Do đó đây được gọi là Inversion of Control.

## Cách dùng

- Không nên new trực tiếp class ở trong 1 class khác sử dụng nó
- Thay vào đó nên khởi tạo class ở trong factory hoặc một nơi thích hợp và class cần sử dụng nó sẽ get object ra thôi. 
- Điều nay sẽ giúp giảm sự phụ thuộc của 2 class giúp nó kết nối với nhau một cách linh hoạt hơn
- Ngoài ra thì các class này có thể test được bằng các mock 