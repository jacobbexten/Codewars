import java.lang.String; 
public class Kata {
  public static String highAndLow(String numbers) {
    String[] nums = numbers.split(" ");
    int low = Integer.parseInt(nums[0]);
    int high = Integer.parseInt(nums[0]);
    
    for (int i = 1; i < nums.length; i++) {
      if (low > Integer.parseInt(nums[i])) {
        low = Integer.parseInt(nums[i]);
      }
      if (high < Integer.parseInt(nums[i])) {
        high = Integer.parseInt(nums[i]);
      }
    }  

    return high + " " + low;
  }
}
