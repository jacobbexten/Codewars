public class HumanReadableTime {
  public static String makeReadable(int seconds) {
    int h = seconds / 3600;
    int m = (seconds % 3600) / 60;
    int s = seconds % 60;
    
    return String.format("%02d:%02d:%02d", h, m, s);
  }
}
