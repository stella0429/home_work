package com.jike;

public class Work01 {
    public static void main(String[] args) {
        int[] data = {9,9,9};
        Work01 work = new Work01();
        work.printArr(data);
        System.out.print("数据处理后：");
        data = work.plusOne(data);
        work.printArr(data);
    }
    public int[] plusOne(int[] digits) {
        int[] temp = new int[digits.length+1];
        int i = digits.length-1;
        int count = 1;
        while(i >= 0){
            if(digits[i]+count == 10){
                temp[i+1] = 0;
            }else {
                temp[i+1] = digits[i]+count;
                if(count > 0){
                    count--;
                }
            }
            i--;
        }
        if(temp[1] == 0){
            temp[0] = 1;
            return temp;
        }else{
            System.arraycopy(temp,1,digits,0,digits.length);
            return digits;
        }
    }

    public void printArr(int[] arr) {
        for(int i = 0;i < arr.length; i++){
            System.out.print(arr[i] + "\t");
        }
        System.out.println();
    }

}
