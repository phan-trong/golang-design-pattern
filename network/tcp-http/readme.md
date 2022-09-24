# HTTP & TCP

- Hypertext Transfer  Protocol (HTTP) và Transmission Control Protocol (TCP) là 2 giao thức truyền tải dữ liệu, nhưng chúng có mục đích riêng và môi quan hệ gần gũi với nhau.

## WHAT IS HTTP:

- HTTP là một giao thức gửi nhận request cho phép người dùng có thể truyền tải dữ liệu trên world wide web(WWW), Giao thức này là phương tiện để sử dụng internet và cung cấp cho 
người dùng cách để tương tác với các tài nguyên HTML trên các máy chủ.
- Đây là giao thức không có trạng thái, có nghĩa là không có thông tin phiên nào từ yêu cầu trước đó được người nhận giữ lại. Tuy nhiên có 1 cách để có thể lưu lại trạng thái thông qua cookie.
- HTTP nằm ở tầng thứ 7(tầng application) của mô hình OSI.

## WHAT IS TCP:

- TCP là giao thức cho phép trao đổi thông tin dữ liệu hoặc message thông qua mạng.
- Nó là một giao thức có trạng thái, và có thể truyền dữ liệu bằng cách chia nhỏ thành các gói tin và đảm bảo phân phối đầu cuối. Nghĩa là nó an toàn và đảm bảo tính toàn vẹn của dữ liệu được gửi qua mạng bật kể số lượng.
- TCP nằm ở lớp thứ 4(the Transport Layer) của mô hình OSI. 
- CÓ nhiều ứng dụng hoạt động dựa trên TCP như HTTP, FTP, email.

## HTTP & TCP, The main differences

- HTTP thường hoạt động ở port 80, đây là cổng sẽ nhận request của user trên những máy chủ. TCP thì không cần yêu cầu điều đó.
- HTTP nhanh hơn TCP, vì nó hoạt động ở tốc độ cao hơn và thực hiện quá trình ngày lập tức. và TCP tương đối chậm hơn.
- HTTP dùng để tìm kiếm trên internet, còn TCP dùng để cho máy tính đích biết ứng dụng nào sẽ nhận dữ liệu và đảm bảo việc phân phối dữ liệu một cách thích hợp.
- TCP chưa thông tin về dữ liệu đã hoặc chưa được nhận, còn HTTP chưa hướng dẫn cụ thể về cách đọc dữ liệu sau khi nhận được.
- TCP quản lý luồng dữ liệu còn HTTP mô tả trong luồng dữ liệu có gì.
- TCP hoạt động 3 chiều trong khi HTTP là giao thức một chiều.

## How to HTTP and TCP work

- thông thường khi user gửi request http thì trước tiên phải thiết lập giao thức TCP vì thế HTTP dựa trên TCP để hoạt động công việc của nó.
