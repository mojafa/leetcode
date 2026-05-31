
class Solution {
    public int minMeetingRooms(List<Interval> intervals) {
        if (intervals == null || intervals.isEmpty()) {
            return 0;
        }

        int n = intervals.size();
        int[] starts = new int[n];
        int[] ends = new int[n];

        for (int i = 0; i < n; i++) {
            Interval in = intervals.get(i);
            starts[i] = in.start; // or in.startTime depending on NeetCode's definition
            ends[i] = in.end;
        }

        Arrays.sort(starts);
        Arrays.sort(ends);

        int rooms = 0, maxRooms = 0;
        int e = 0;

        for (int s = 0; s < n; s++) {
            while (e < n && ends[e] <= starts[s]) {
                rooms--;
                e++;
            }
            rooms++;
            maxRooms = Math.max(maxRooms, rooms);
        }

        return maxRooms;
    }
}
