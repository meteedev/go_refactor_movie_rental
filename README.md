# คำอธิบายโปรแกรม

โปรแกรมนี้ใช้คำนวณค่าเช่าหนังให้กับลูกค้า โดยการคิดค่าบริการจะขึ้นอยู่กับระยะเวลาของการเช่าและชนิดของหนัง
โดยหนังมีทั้งหมด 3 ประเภท คือ หนังปกติ(regular) หนังเด็ก(children's) และหนังใหม่(new releases)
และโปรแกรมยังสามารถคำนวณแต้มสำหรับผู้เช่าบ่อย ซึ่งหากเป็นหนังใหม่ก็จะได้แต้มเป็น 2 เท่า

# เช่าหนัง

ความต้องการ:

ตอนนี้ `statement` แสดงผลเป็นรูปแบบตัวอักษรปกติ

```
Rental Record for AnuchitO
 	Kingsman	3.5
	Iron Man	2.5
Amount owed is 6.0
You earned 2 frequent renter points
```

ฝ่ายขายอยากได้การแสดงผลเป็นรูปแบบ HTML :

```
<h1>Rental Record for <em>AnuchitO</em></h1>
<table>
  <tr><td>Kingsman</td><td>3.5</td></tr>
  <tr><td>Iron Man</td><td>2.5</td></tr>
</table>
<p>Amount owed is <em>6.0</em></p>
<p>You earned <em>2</em> frequent renter points</p>
```
# Refactor step
1. rename variable 
2. remove pointer
3. fix New method return direct struct
4. remove "get" method get not start with get
5. constant iota
6. fix format statement print format
7. extract method amountFor from  statement
8. rename method  amountFor to charge 
9. receiver rental with charge method
10. extract method getPoints from statement
11. split loop statment to getTotalAmount  , getTotalPoints
12. push dependency getTotalAmount , getTotalPoints
13. create struct Record , Charge for renderPlainText
14. separate renderPlainText for presentation only
15. remove switch Charge to use polymorp Charge