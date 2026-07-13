# AI Algorithm Implementation Suite

Dự án này là bộ sưu tập các thuật toán cơ bản trong Trí tuệ nhân tạo, bao gồm các thuật toán tìm đường (Path Finding) và các thuật toán suy diễn logic mệnh đề (Sentential Logic).

## 1. Cấu trúc dữ liệu cấu hình

Hệ thống sử dụng hai file cấu hình chính nằm trong thư mục `data/`:

### a. `data/graph.json` (Cấu hình đồ thị)

Sử dụng định dạng JSON để mô tả đồ thị.

* **Cấu trúc:** Mỗi nút là một key, giá trị bao gồm:
* `path`: Map các nút con và trọng số của cạnh.
* `isAnd`: Danh sách các nút con kết nối bằng toán tử AND (nếu để trống, mặc định là OR).
* `heuristic`: Giá trị ước lượng cho các thuật toán tìm kiếm thông minh (A*, Greedy).


* **Ví dụ:**

```json
{
  "A": {"path": {"B": 1, "C": 1}, "isAnd": [], "heuristic": 0},
  "B": {"path": {"D": 1, "E": 1}, "isAnd": ["D", "E"], "heuristic": 0}
}

```

### b. `data/logic.md` (Cấu hình logic)

Sử dụng cú pháp gần giống LaTeX để định nghĩa các mệnh đề logic.

* **Quy tắc:**
* Sử dụng: `\land` (AND), `\lor` (OR), `\neg` (NOT), `\rightarrow` (Implies).
* Dấu `\implies` dùng để ngăn cách giữa tập tiền đề (Premise) và kết luận (Conclusion).
* Các mệnh đề đơn lẻ không có dấu `\implies` sẽ được hệ thống hiểu là các sự thật (Facts).


* **Ví dụ:**

```markdown
a \land c
a \land b \rightarrow f
(d \lor b) \land f \rightarrow i
\implies i

```

## 2. Các thuật toán hỗ trợ

### Tìm đường (Path Finding)

* `dfs`: Tìm kiếm chiều sâu.
* `bfs`: Tìm kiếm chiều rộng.
* `min`: Greedy Search.
* `A*`: Tìm kiếm A*.
* `hill`: Hill Climbing.

### Logic mệnh đề (Sentential Logic)

* `fc`: Forward Chaining (Suy diễn tiến).
* `bc`: Backward Chaining (Suy diễn lùi).
* `wa`: Vương Hạo (Wang's Algorithm).
* `r`: Robinson (Resolution).

## 3. Hướng dẫn chạy chương trình

### Yêu cầu

* Ngôn ngữ: Go (phiên bản 1.25 trở lên).

### Các bước thực hiện

1. **Clone dự án** và di chuyển vào thư mục gốc.
2. **Chạy chương trình:**
Sử dụng lệnh `go run` từ thư mục gốc của dự án:
```bash
go run ./cli

```


3. **Sử dụng:**
Sau khi hệ thống khởi động, bạn sẽ thấy dấu nhắc lệnh `>`. Nhập lệnh theo cú pháp:
```text
<algo> <from> <to>

```


* Ví dụ tìm đường: `dfs A B`
* Ví dụ chạy thuật toán logic: `fc` hoặc `r` (không cần from/to).


4. **Thoát:** Nhập `q` hoặc `quit` để thoát chương trình.

---

### Mẹo nhỏ cho người dùng

* Các file cấu hình `data/graph.json` và `data/logic.md` có thể được chỉnh sửa trực tiếp, chương trình sẽ tự động nạp lại dữ liệu khi khởi động lại.
---