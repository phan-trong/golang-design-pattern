# Design Pattern Builder

## Diễn giải

- Builder là một công cụ để xây dựng những đối tượng phức tạp, lần lượt từng bước.
- Những đôi tượng có kiểu khác nhau, biểu diễn khác nhau nhưng có chung các bước để xây dựng

## Vấn đề

- Thử tưởng tượng có một đối tượng phức tạp với nhiều biến thể con, nhiều field và những đối tượng lồng nhau. 
Như vậy code sẽ khiến hàm khởi tạo có nhiều tham số hoặc phải check ở nhiều nơi, 
Ví Dụ:

- Để xây dựng một căn nhà sẽ cần đi qua các bước xây dựng móng, 4 bức tường, và mái.
Nhưng nếu muốn xây dựng một ngồi nhà rộng hơn có nhiều tiện nghi hơn thì cách đơn giản nhất là viết 1 sub class thừa kế từ class cha.
Tuy nhiên như vậy chúng ta sẽ phải tạo ra nhiều sub class và nhiều params.

- Có một giải pháp khác là tạo một lớp nhiều với nhiều tham số, với class nào cần dùng tham số nào thì sẽ truyền giá trị vào với tham số đó
Tuy nhiên cách trên cũng sẽ làm xấu đi hàm khởi tạo và dư thừa nhiều tham số không cần dùng đến.

## Giải Pháp

- Xây dựng một khung builder, tách những đối tượng ra khỏi lớp cha và chia nhỏ, tạo ra các đối tượng riêng biệt
- Mô hình này tổ chức khởi tạo đối tượng bằng cách thành lập một tập hợp các bước. Để tạo đối tượng, cần phải gọi một loạt các bước cần thiết để tạo ra đối tượng cụ thể.
- Các bước để tạo đối tượng sẽ được gọi ở lớp director. Lớp này không bắt buộc phải có tuy nhiên thì đây là một nơi tốt để chạy các bước xây dựng đối tượng và tái sử dụng các bước đó.


### Nguồn: 
https://refactoring.guru/design-patterns/builder
https://refactoring.guru/design-patterns/builder/go/example#example-0