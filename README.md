WB Tech: level # 1 (Golang)
��� ������ �������
� �������� ������� ������ ������� � ������ ���. ���� ������� � ���� ���� �
������ ������������������ �����. ������ ������� ��� ������������� �������
���� ���������.
����������� � �������������� ������������� ����� ���������� ��������,
����������� ��������� ��������� � �.�. � �.�.
�������� �������� ������ � ������ ��������� ���� ��� ��������. ��������� ������
����� ������ ����������� ���������, � ���� ������ ��������� ��������
����������� ��������� ���������� ���������.
����� �������� �������, ��� �� ������� �����, ��� � �� �� �������. ���������
������� � ������������������ ���� ������� � �������� ������������ ������ ��
������� ������������� Wildberries.
�������
1. ���� ��������� Human (� ������������ ������� ����� � �������).
����������� ����������� ������� � ��������� Action �� ������������ ���������
Human (������ ������������).
2. �������� ���������, ������� ����������� ���������� �������� ��������� �����
������ �� ������� (2,4,6,8,10) � ������� �� �������� � stdout.
3. ���� ������������������ �����: 2,4,6,8,10. ����� ����� ��
���������(2
2+3
2+4
2�.) � �������������� ������������ ����������.
4. ����������� ���������� ������ ������ � ����� (������� �����). �����������
����� �� N ��������, ������� ������ ������������ ������ �� ������ �
������� � stdout. ���������� ����������� ������ ���������� �������� ���
������.
��������� ������ ����������� �� ������� Ctrl+C. ������� � ����������
������ ���������� ������ ���� ��������.
5. ����������� ���������, ������� ����� ��������������� ���������� �������� �
�����, � � ������ ������� ������ � ������. �� ��������� N ������ ���������
������ �����������.
6. ����������� ��� ��������� ������� ��������� ���������� ��������.
7. ����������� ������������ ������ ������ � map.
8. ���� ���������� int64. ����������� ��������� ������� ������������� i-� ��� �
1 ��� 0.
9. ����������� �������� �����. ���� ��� ������: � ������ ������� ����� (x) ��
�������, �� ������ � ��������� �������� x*2, ����� ���� ������ �� �������
������ ������ ���������� � stdout.
10. ���� ������������������ ������������� ���������: -25.4, -27.0 13.0, 19.0,
15.5, 24.5, -21.0, 32.5. ���������� ������ �������� � ������ � ����� � 10
��������. ������������������ � �������������� �� �����.
������: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
11. ����������� ����������� ���� ��������������� ��������.
12. ������� ������������������ ����� - (cat, cat, dog, cat, tree) ������� ��� ���
����������� ���������.
13. �������� ������� ��� ����� ��� �������� ��������� ����������.
14. ����������� ���������, ������� � �������� �������� ���������� ���
����������: int, string, bool, channel �� ���������� ���� interface{}.
15. � ����� ���������� ������������ ����� �������� ������ �������� ����, � ���
��� ���������? ��������� ���������� ������ ����������.
var justString string
func someFunc() {
v := createHugeString(1 << 10)
justString = v[:100]
}
func main() {
someFunc()
}
16. ����������� ������� ���������� ������� (quicksort) ����������� ��������
�����.
17. ����������� �������� ����� ����������� �������� �����.
18. ����������� ���������-�������, ������� ����� ������������������ �
������������ �����. �� ���������� ��������� ������ �������� ��������
�������� ��������.
19. ����������� ���������, ������� �������������� ���������� �� ��� ������
(��������: ��������� � ��������). ������� ����� ���� unicode.
20. ����������� ���������, ������� �������������� ����� � ������.
������: �snow dog sun � sun dog snow�.
21. ����������� ������� �������� �� ����� �������.
22. ����������� ���������, ������� �����������, �����, ����������, �������� ���
�������� ���������� a,b, �������� ������� > 2^20.
23. ������� i-�� ������� �� ������.
24. ����������� ��������� ���������� ���������� ����� ����� �������, �������
������������ � ���� ��������� Point � ������������������ ����������� x,y �
�������������.
25. ����������� ����������� ������� sleep.
26. ����������� ���������, ������� ���������, ��� ��� ������� � ������
���������� (true � ���� ����������, false etc). ������� �������� ������ ����
�������������������.
��������:
abcd � true
abCdefAaf � false
aabcd � false
������ �������
1. ����� ����� ����������� ������ ������������ �����?
2. ��� ����� ����������, ��� ��� ����������� � Go?
3. ��� ���������� RWMutex �� Mutex?
4. ��� ���������� ���������������� � �� ���������������� ������?
5. ����� ������ � ��������� struct{}{}?
6. ���� �� � Go ���������� ������� ��� ����������?
7. � ����� ������������������ ����� �������� �������� map[int]int?
������:
m[0]=1
m[1]=124
m[2]=281
8. � ��� ������� make � new?
9. ������� ���������� �������� ������ ���������� ���� slice ��� map?
10. ��� ������� ������ ��������� � ������?
func update(p *int) {
b := 2
p = &b
}
func main() {
var (
a = 1
p = &a
)
fmt.Println(*p)
update(p)
fmt.Println(*p)
}
11. ��� ������� ������ ��������� � ������?
func main() {
wg := sync.WaitGroup{}
for i := 0; i < 5; i++ {
wg.Add(1)
go func(wg sync.WaitGroup, i int) {
fmt.Println(i)
wg.Done()
}(wg, i)
}
wg.Wait()
fmt.Println("exit")
}
12. ��� ������� ������ ��������� � ������?
func main() {
n := 0
if true {
n := 1
n++
}
fmt.Println(n)
}
13. ��� ������� ������ ��������� � ������?
func someAction(v []int8, b int8) {
v[0] = 100
v = append(v, b)
}
func main() {
var a = []int8{1, 2, 3, 4, 5}
someAction(a, 6)
fmt.Println(a)
}
14. ��� ������� ������ ��������� � ������?
func main() {
slice := []string{"a", "a"}
func(slice []string) {
slice = append(slice, "a")
slice[0] = "b"
slice[1] = "b"
fmt.Print(slice)
}(slice)
fmt.Print(slice)
}