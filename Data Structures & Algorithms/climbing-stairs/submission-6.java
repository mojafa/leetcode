class Solution {
    public int climbStairs(int n) {
        if (n<=2){
            return n;
        }

        int a = 1; // ways for step 1 behind (i-2)
        int b = 2; // ways for step 1 behind (i-1)

        for (int i =3; i<=n;i++){
            int c = a+b;
            a =b;
            b=c;
        }
        return b;


    }
}
