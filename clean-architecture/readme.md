# Clean Architecture

![image description cleab architecture](./CleanArchitecture.jpg)

- Sơ đồ trên là kiến trúc của clean architecture
- kiến trúc này gồm 4 layer chính:
    + Entities
    + Use Cases
    + Interface Adapters
    + Frameworks and Drivers

## The Dependency Rule (Quy tắc phụ thuộc)

- Ở vòng tròn đồng tâm trên ta thấy đây nó biểu diễn những vùng khác nhau. Tuy nhiên, các vòng tròn bên ngoài là cơ chế, các vòng tròn bên trong là chính sách.
- Quy tắc này nói răng các phụ thuộc code chỉ nên hướng vào bên trong vòng tròn.
- Vòng tròn bên trong không thể phụ thuộc vào cái gì đó ở vòng tròn bên ngoài.
- Đặc biệt, khi khai báo tên không được đề cập bởi code ở vòng tròn bên trong. Bao gồm cả tên hàm tên lớp, tên biến hoặc bất cứ thực thể nào khác.
- Tương tự, các dữ liệu bên ngoài không được sử dụng ở vòng tròn bên trong.

## Entities

- Entities đóng gói toàn bộ business rule, một entites có thể là một object với phương thức, hoặc có thể là một tập hợp cấu trúc hoặc chức năng(funtions), điều đó không thành vấn đè miễn là có thể sử dụng bởi nhiều class khác.
- Nếu bạn không có enterprise, và chỉ đang code 1 ứng dụng duy nhất thì entites là những business object của ứng dụng.
- Những entites thường được đống gói lại, và ít khi bị thay đổi nhất, không có hoạt động thay đổi nào sẽ ảnh hưởng đến lớp này.

## Use Cases

- Ở layer này nó bao gồm những business rule cụ thể, nó được đóng gói và implement tất cả use case của hệ thống. Những use case này điều phối luồng dữ liệu đến và đi từ entities.
- Những thay đổi ở lớp này không ảnh hướng đến lớp entities. Và lớp này cũng sẽ không bị ảnh hưởng bởi những thay đổi ở lớp bên ngoài như database, UI hoặc bất cứ common framework nào. Layer này là độc lập.
- Lớp này chỉ bị sửa đôi khi user mong muốn sửa use case.

## Interface Adapters

- Đây là layer tập hợp các bộ chuyển đổi format dữ liệu thuận tiện nhất cho lớp use cases và lớp entities.
- Chính lớp này sẽ chưa kiến trúc MVC của GUI. The presenters, views, Controllers sẽ nằm ở đây.
- Ví dụ ở lớp này database là SQL thì tất cả SQL phải giới hạn ở lớp này.

## Frameworks and Drivers

- Ở layer này thường code ít và chỉ là nơi giao tiếp sử dụng các lớp ở bên trong.

## Only Four Circles?

- Không phải lúc nào cũng chỉ có 4 vòng tròn ở trên sơ đồ, không có quy tắc nói luôn luôn chỉ có 4 vòng tròn này. Tuy nhiên nguyên tắc dependency luôn luôn được áp dụng. Nghĩa là những mũi tên hướng vào bên trong, càng vào bên trong càng trừu tượng.
- Vòng bên trong cùng là vòng chung nhất.

Nguồn:
- https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html
