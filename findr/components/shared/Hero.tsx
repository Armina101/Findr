import { jobType } from "@/types/jobTypes";
import SearchForm from "./SearchForm";

interface HeroProps {
  jobs: jobType[];
}

const Hero = ({ jobs }: HeroProps) => {
  return (
    <div
      id="home"
      className="bg-[url('/hero.jpg')] h-screen relative w-full bg-cover mt-[-70px]"
    >
      <div className="flex flex-col h-full items-center justify-center pt-[82px] gap-20 w-[90%] mx-auto max-w-[1450px]">
        <div className="text-center flex flex-col gap-3">
          <h1 className="text-6xl font-extrabold text-black">
            Get The{" "}
            <span className="text-blue-600 font-extrabold">
              Ideal Placement
            </span>
            , You Deserve
          </h1>
          <span className="font-thin">
            Placement Opportunities listed here! Your ideal placement is
            waiting
          </span>
        </div>
        <SearchForm jobs={jobs} />
      </div>
    </div>
  );
};

export default Hero;