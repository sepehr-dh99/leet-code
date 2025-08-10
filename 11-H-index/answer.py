class Solution:
    def hIndex(self, citations: list[int]) -> int:
        citation = 0
        citations.sort(reverse=True)

        for i in range(len(citations)):
            if citations[i] > i:
                citation += 1

        return citation