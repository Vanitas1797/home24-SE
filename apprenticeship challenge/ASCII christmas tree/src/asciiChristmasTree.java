
/**
 * @author Maik Stegemann
 *
 */
public class asciiChristmasTree {
	
	private static int height;
	public static String star;
	
	/**
	 * @param height size of the tree as an integer
	 * @param star decision whether yes or no
	 */
	private static void buildTree(int height) {
		int temp = 0;
		System.out.println();
		for(int i = height; i > 0; i--) {
			for(int j = i - 1; j > 0; j--) {
				System.out.print(" ");
			}
			if(temp == 0) {
				if(star.equals("yes")) {
					System.out.print("*");
					i++;
				}
				temp++;
			}
			for(int k = (height - i) * 2 + 1; k > 0; k--) {
				System.out.print("X");
			}
			System.out.println();
		}
		for(int l = height - 1; l > 0; l--) {
			System.out.print(" ");
		}
		System.out.print("I");
		System.out.println();
	}

	/**
	 * @param args[0] convertion from string to integer that is inserted
	 * @param args[1] "yes" or anything else except empty as options for the star on top of the tree
	 */
	public static void main(String[] args) {
		height = Integer.parseInt(args[0]);
		star = args[1];
		buildTree(height);
	}
}
